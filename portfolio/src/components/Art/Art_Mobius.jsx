import NavBar from "../NavBar";
function Art_Mobius() {
  const body = document.body;
  body.className = "default-bg";
  const items = [
    {
      id: 2,
      title: "Mercury Lounge Poster",
              image: `/mobius/Mercury_Lounge.jpg`,
      comment: "",
    },
    {
      id: 3,
      title: "Tractor Tavern Poster",
              image: `/mobius/Tractor_Tavern.jpg`,
      comment: "",
    },
    {
      id: 4,
      title: "Grey Eagle Poster",
              image: `/mobius/Grey_Eagle.jpg`,
      comment: "",
    },
    {
      id: 1,
      title: "QR Code",
              image: `/mobius/QR.jpg`,
      comment: "",
    },
    {
      id: 5,
      title: "Promotional Image 1",
              image: `/mobius/elevator.png`,
      comment: "",
    },
    {
      id: 6,
      title: "Promotional Image 4",
              image: `/mobius/climb.png`,
      comment: "",
    },
  ];

  return (
    <div className="sketchOuterContainer">
      <div className="navBarDiv">
        <NavBar scrollText="HAVE YOU HEARD THIS SONG? HAVE YOU HEARD THIS SONG? HAVE YOU HEARD THIS SONG? HAVE YOU HEARD THIS SONG?" />
      </div>
      <div className="sketchTitle">
        <img
          className="mobiusArtPageTitle"
          src="/mobius/Logo2.jpg"
        />
        <img
          className="mobiusFooter"
          src="/mobius/croppedbeatles.png"
        />
        <p>
          MOBIUS LINE is an ARG (Augmented Reality Game) that I co-created with
          Sumner Badzik. This project involved creating branding and marketing
          for our fictitious band MOBIUS LINE, as well as constructing a
          narrative which is scattered across multiple social media platforms
          and accounts. Originally, posters (seen below) were hung around the
          neighborbood, linking to the start of the game.On this page I have
          included some of the graphic design work I did for the project, as
          well as some of our band photos, taken by Andrew Yurovchak. I hope
          you'll take a little time to explore the game{" "}
          <a target="_blank" href="https://m.youtube.com/watch?v=VJmVcI-avlE">
            here
          </a>
          .
        </p>
      </div>

      <div className="sketchContainer">
        {items.map((project) => (
          <div className="sketchDiv" key={project.id}>
            <img className="sketchImage" src={project.image} />
            <h2>{project.title}</h2>
            <p>{project.comment}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
export default Art_Mobius;
