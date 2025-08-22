package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/mail"
    "os"
    "os/exec"
    "path/filepath"
    "sort"
    "strings"
    "time"

    tea "github.com/charmbracelet/bubbletea"
)

type Email struct {
    From     string
    To       string
    Subject  string
    Date     string
    Body     string
    Filename string
}

type state int

const (
    mainMenuState state = iota
    emailListState
    emailViewState
    composeState
)

type model struct {
    state       state
    cursor      int
    emails      []*Email
    currentEmail *Email
    
    // Compose fields
    composeTo      string
    composeSubject string
    composeBody    string
    composeField   int // 0=to, 1=subject, 2=body
    
    // UI
    width  int
    height int
}

func parseEmail(filename string) (*Email, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    msg, err := mail.ReadMessage(file)
    if err != nil {
        return nil, err
    }

    body, err := io.ReadAll(msg.Body)
    if err != nil {
        return nil, err
    }

    email := &Email{
        From:     msg.Header.Get("From"),
        To:       msg.Header.Get("To"),
        Subject:  msg.Header.Get("Subject"),
        Date:     msg.Header.Get("Date"),
        Body:     strings.TrimSpace(string(body)),
        Filename: filename,
    }

    return email, nil
}

func getEmails() ([]*Email, error) {
    maildir := filepath.Join(os.Getenv("HOME"), "Maildir", "new")
    files, err := ioutil.ReadDir(maildir)
    if err != nil {
        return nil, err
    }
    
    var emails []*Email
    for _, file := range files {
        emailPath := filepath.Join(maildir, file.Name())
        email, err := parseEmail(emailPath)
        if err == nil {
            emails = append(emails, email)
        }
    }
    
    sort.Slice(emails, func(i, j int) bool {
        return emails[i].Date > emails[j].Date
    })
    
    return emails, nil
}

func sendEmail(to, subject, body string) error {
    from := fmt.Sprintf("%s@nicholasmirigliani.com", os.Getenv("USER"))
    date := time.Now().Format(time.RFC1123Z)
    
    emailContent := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nDate: %s\n\n%s\n",
        from, to, subject, date, body)
    
    // Send using sendmail
    cmd := exec.Command("sendmail", to)
    cmd.Stdin = strings.NewReader(emailContent)
    err := cmd.Run()
    if err != nil {
        return err
    }
    
    // Also save to sent folder
    sentDir := filepath.Join(os.Getenv("HOME"), "Maildir", ".Sent")
    err = os.MkdirAll(sentDir, 0755)
    if err != nil {
        return err
    }
    
    timestamp := time.Now().Unix()
    filename := filepath.Join(sentDir, fmt.Sprintf("%d.eml", timestamp))
    
    return ioutil.WriteFile(filename, []byte(emailContent), 0644)
}

func initialModel() model {
    return model{
        state: mainMenuState,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        return m, nil
    
    case tea.KeyMsg:
        switch m.state {
        case mainMenuState:
            return m.updateMainMenu(msg)
        case emailListState:
            return m.updateEmailList(msg)
        case emailViewState:
            return m.updateEmailView(msg)
        case composeState:
            return m.updateCompose(msg)
        }
    }
    
    return m, nil
}

func (m model) updateMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "q", "ctrl+c":
        return m, tea.Quit
    case "up", "k":
        if m.cursor > 0 {
            m.cursor--
        }
    case "down", "j":
        if m.cursor < 2 {
            m.cursor++
        }
    case "enter", " ":
        switch m.cursor {
        case 0: // View Emails
            emails, _ := getEmails()
            m.emails = emails
            m.state = emailListState
            m.cursor = 0
        case 1: // Send Email
            m.state = composeState
            m.cursor = 0
            m.composeField = 0
            m.composeTo = ""
            m.composeSubject = ""
            m.composeBody = ""
        case 2: // Quit
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) updateEmailList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "q", "ctrl+c":
        return m, tea.Quit
    case "b", "esc":
        m.state = mainMenuState
        m.cursor = 0
    case "up", "k":
        if m.cursor > 0 {
            m.cursor--
        }
    case "down", "j":
        if m.cursor < len(m.emails)-1 {
            m.cursor++
        }
    case "enter", " ":
        if len(m.emails) > 0 && m.cursor < len(m.emails) {
            m.currentEmail = m.emails[m.cursor]
            m.state = emailViewState
        }
    case "d":
        if len(m.emails) > 0 && m.cursor < len(m.emails) {
            email := m.emails[m.cursor]
            os.Remove(email.Filename)
            // Refresh email list
            emails, _ := getEmails()
            m.emails = emails
            if m.cursor >= len(m.emails) && len(m.emails) > 0 {
                m.cursor = len(m.emails) - 1
            }
        }
    }
    return m, nil
}

func (m model) updateEmailView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "q", "ctrl+c":
        return m, tea.Quit
    case "b", "esc":
        m.state = emailListState
    case "d":
        if m.currentEmail != nil {
            os.Remove(m.currentEmail.Filename)
            m.state = emailListState
            // Refresh email list
            emails, _ := getEmails()
            m.emails = emails
            if m.cursor >= len(m.emails) && len(m.emails) > 0 {
                m.cursor = len(m.emails) - 1
            }
        }
    }
    return m, nil
}

func (m model) updateCompose(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "ctrl+c":
        return m, tea.Quit
    case "esc":
        m.state = mainMenuState
        m.cursor = 0
    case "tab":
        if m.composeField < 2 {
            m.composeField++
        }
    case "shift+tab":
        if m.composeField > 0 {
            m.composeField--
        }
    case "ctrl+s":
        err := sendEmail(m.composeTo, m.composeSubject, m.composeBody)
        if err == nil {
            m.state = mainMenuState
            m.cursor = 0
        }
    case "enter":
        if m.composeField == 2 { // Body field
            m.composeBody += "\n"
        } else {
            if m.composeField < 2 {
                m.composeField++
            }
        }
    case "backspace":
        switch m.composeField {
        case 0:
            if len(m.composeTo) > 0 {
                m.composeTo = m.composeTo[:len(m.composeTo)-1]
            }
        case 1:
            if len(m.composeSubject) > 0 {
                m.composeSubject = m.composeSubject[:len(m.composeSubject)-1]
            }
        case 2:
            if len(m.composeBody) > 0 {
                m.composeBody = m.composeBody[:len(m.composeBody)-1]
            }
        }
    default:
        char := msg.String()
        if len(char) == 1 {
            switch m.composeField {
            case 0:
                m.composeTo += char
            case 1:
                m.composeSubject += char
            case 2:
                m.composeBody += char
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    switch m.state {
    case mainMenuState:
        return m.viewMainMenu()
    case emailListState:
        return m.viewEmailList()
    case emailViewState:
        return m.viewEmail()
    case composeState:
        return m.viewCompose()
    }
    return ""
}

func (m model) viewMainMenu() string {
    username := os.Getenv("USER")
    emailLine := fmt.Sprintf("   Logged in as: %s@nicholasmirigliani.com   ", username)
    titleLine := "          Email Client v1.0           "
    
    width := len(emailLine)
    if len(titleLine) > width {
        width = len(titleLine)
    }
    
    var s strings.Builder
    
    // Banner
    s.WriteString("╔")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╗\n")
    
    // Title line
    padding := (width - len(titleLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(titleLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(titleLine)))
    s.WriteString("║\n")
    
    // Email line
    padding = (width - len(emailLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(emailLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(emailLine)))
    s.WriteString("║\n")
    
    s.WriteString("╚")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╝\n\n")
    
    // Menu
    s.WriteString("┌─────────────────────┐\n")
    s.WriteString("│     MAIN MENU       │\n")
    s.WriteString("├─────────────────────┤\n")
    
    options := []string{"View Emails", "Send Email", "Quit"}
    for i, option := range options {
        if i == m.cursor {
            s.WriteString(fmt.Sprintf("│ ► %d. %-14s │\n", i+1, option))
        } else {
            s.WriteString(fmt.Sprintf("│   %d. %-14s │\n", i+1, option))
        }
    }
    
    s.WriteString("└─────────────────────┘\n\n")
    s.WriteString("Use ↑/↓ arrows to navigate, Enter to select, q to quit")
    
    return s.String()
}

func (m model) viewEmailList() string {
    var s strings.Builder
    
    username := os.Getenv("USER")
    emailLine := fmt.Sprintf("   Logged in as: %s@nicholasmirigliani.com   ", username)
    titleLine := "          Email Client v1.0           "
    
    width := len(emailLine)
    if len(titleLine) > width {
        width = len(titleLine)
    }
    
    // Banner
    s.WriteString("╔")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╗\n")
    
    // Title line
    padding := (width - len(titleLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(titleLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(titleLine)))
    s.WriteString("║\n")
    
    // Email line
    padding = (width - len(emailLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(emailLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(emailLine)))
    s.WriteString("║\n")
    
    s.WriteString("╚")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╝\n\n")
    
    if len(m.emails) == 0 {
        s.WriteString("📭 No new emails found.\n\n")
        s.WriteString("Press 'b' to go back to main menu")
        return s.String()
    }
    
    s.WriteString(fmt.Sprintf("📧 Found %d new emails:\n\n", len(m.emails)))
    s.WriteString("┌────┬──────────────────────┬─────────────────────────────────────────┐\n")
    s.WriteString("│ #  │ From                 │ Subject                                 │\n")
    s.WriteString("├────┼──────────────────────┼─────────────────────────────────────────┤\n")
    
    for i, email := range m.emails {
        from := email.From
        if len(from) > 20 {
            from = from[:17] + "..."
        }
        subject := email.Subject
        if len(subject) > 39 {
            subject = subject[:36] + "..."
        }
        
        if i == m.cursor {
            s.WriteString(fmt.Sprintf("│►%2d │ %-20s │ %-39s │\n", i+1, from, subject))
        } else {
            s.WriteString(fmt.Sprintf("│ %2d │ %-20s │ %-39s │\n", i+1, from, subject))
        }
    }
    s.WriteString("└────┴──────────────────────┴─────────────────────────────────────────┘\n\n")
    s.WriteString("↑/↓ to navigate, Enter to read, 'd' to delete, 'b' to go back")
    
    return s.String()
}

func (m model) viewEmail() string {
    if m.currentEmail == nil {
        return "No email selected"
    }
    
    var s strings.Builder
    
    username := os.Getenv("USER")
    emailLine := fmt.Sprintf("   Logged in as: %s@nicholasmirigliani.com   ", username)
    titleLine := "          Email Client v1.0           "
    
    width := len(emailLine)
    if len(titleLine) > width {
        width = len(titleLine)
    }
    
    // Banner
    s.WriteString("╔")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╗\n")
    
    // Title line
    padding := (width - len(titleLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(titleLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(titleLine)))
    s.WriteString("║\n")
    
    // Email line
    padding = (width - len(emailLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(emailLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(emailLine)))
    s.WriteString("║\n")
    
    s.WriteString("╚")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╝\n\n")
    
    s.WriteString("╔════════════════════════════════════════════════════════════════════╗\n")
    s.WriteString("║                            EMAIL DETAILS                          ║\n")
    s.WriteString("╚════════════════════════════════════════════════════════════════════╝\n")
    s.WriteString(fmt.Sprintf("From:    %s\n", m.currentEmail.From))
    s.WriteString(fmt.Sprintf("To:      %s\n", m.currentEmail.To))
    s.WriteString(fmt.Sprintf("Subject: %s\n", m.currentEmail.Subject))
    s.WriteString(fmt.Sprintf("Date:    %s\n", m.currentEmail.Date))
    s.WriteString("─────────────────────────────────────────────────────────────────────\n")
    s.WriteString(m.currentEmail.Body)
    s.WriteString("\n─────────────────────────────────────────────────────────────────────\n\n")
    s.WriteString("'d' to delete, 'b' to go back")
    
    return s.String()
}

func (m model) viewCompose() string {
    var s strings.Builder
    
    username := os.Getenv("USER")
    emailLine := fmt.Sprintf("   Logged in as: %s@nicholasmirigliani.com   ", username)
    titleLine := "          Email Client v1.0           "
    
    width := len(emailLine)
    if len(titleLine) > width {
        width = len(titleLine)
    }
    
    // Banner
    s.WriteString("╔")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╗\n")
    
    // Title line
    padding := (width - len(titleLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(titleLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(titleLine)))
    s.WriteString("║\n")
    
    // Email line
    padding = (width - len(emailLine)) / 2
    s.WriteString("║")
    s.WriteString(strings.Repeat(" ", padding))
    s.WriteString(emailLine)
    s.WriteString(strings.Repeat(" ", width-padding-len(emailLine)))
    s.WriteString("║\n")
    
    s.WriteString("╚")
    s.WriteString(strings.Repeat("═", width))
    s.WriteString("╝\n\n")
    
    s.WriteString("╔════════════════════════════════════════════════════════════════════╗\n")
    s.WriteString("║                           COMPOSE EMAIL                           ║\n")
    s.WriteString("╚════════════════════════════════════════════════════════════════════╝\n")
    
    // To field
    if m.composeField == 0 {
        s.WriteString(fmt.Sprintf("To: %s█\n", m.composeTo))
    } else {
        s.WriteString(fmt.Sprintf("To: %s\n", m.composeTo))
    }
    
    // Subject field
    if m.composeField == 1 {
        s.WriteString(fmt.Sprintf("Subject: %s█\n", m.composeSubject))
    } else {
        s.WriteString(fmt.Sprintf("Subject: %s\n", m.composeSubject))
    }
    
    s.WriteString("Body:\n")
    if m.composeField == 2 {
        s.WriteString(fmt.Sprintf("%s█\n", m.composeBody))
    } else {
        s.WriteString(fmt.Sprintf("%s\n", m.composeBody))
    }
    
    s.WriteString("\nTab/Shift+Tab to navigate fields, Ctrl+S to send, Esc to cancel")
    
    return s.String()
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}