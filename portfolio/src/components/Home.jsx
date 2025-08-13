import NavBar from "./NavBar";
function Home() {
  const body = document.body;
  body.className = "home-bg";
  return (
    <div className="homeContainer">
      <div>
        <NavBar scrollText="WHO IS THIS GUY? WHAT DOES HE WANT FROM YOU? WHAT IS HE SMILING ABOUT? WHO IS THIS GUY? WHAT DOES HE WANT FROM YOU? WHAT IS HE SMILING ABOUT?" />
      </div>
      <div className="homeInnerContainer1">
        <div className="homeTextDiv">
          <h2>About</h2>
          <p className="homeText">
            I am a Computer Science student at Stevens Institute of Technology,
            pursuing a minor in Visual Arts & Technology. My work consists
            primarily of ballpoint pen drawings, but spans a wide array of
            mediums. I have done work in web design, graphic design,
            photography, videography, collage, poetry, and creative writing (and
            even a little music composition.) <br />
            This website is a platform for me to share all of my creative
            endeavors, thanks for checking it out!
            <br />
            Inquiries can be sent to{" "}
            <a href="mailto:nicholasmirigliani@gmail.com">
              nicholasmirigliani@gmail.com
            </a>
            .
          </p>
        </div>
        <img
          className="homeImage"
          src="/images/me.JPEG"
        />
      </div>
      <div className="homeInnerContainer2">
        <h2>Links</h2>
        <a target="_blank" href="https://www.youtube.com/@nobbybop">
          YouTube
        </a>
        <a target="_blank" href="https://www.instagram.com/nick.mirig/">
          Instagram
        </a>
        <a
          target="_blank"
          href="https://www.linkedin.com/in/nicholas-mirigliani/"
        >
          LinkedIn
        </a>
        <a target="_blank" href="https://github.com/NobbyBop">
          GitHub
        </a>
        <a target="_blank" href="/Resume.pdf">
          Résumé
        </a>
      </div>
    </div>
  );
}
export default Home;
