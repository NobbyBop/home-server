import NavBar from "../NavBar";
function Art_Worldly() {
  const body = document.body;
  body.className = "default-bg";
  const items = [
    {
      id: 1,
      title: "Cover",
      image: `/worldly/1.JPG`,
      comment: "",
    },
    {
      id: 2,
      title: "Page 1",
      image: `/worldly/2.JPG`,
      comment: "",
    },
    {
      id: 3,
      title: "Page 2",
      image: `/worldly/3.JPG`,
      comment: "",
    },
    {
      id: 4,
      title: "Page 3",
      image: `/worldly/4.JPG`,
      comment: "",
    },
    {
      id: 5,
      title: "Page 4",
      image: `/worldly/5.JPG`,
      comment: "",
    },
    {
      id: 6,
      title: "Page 5",
      image: `/worldly/6.JPG`,
      comment: "",
    },
    {
      id: 7,
      title: "Page 6",
      image: `/worldly/7.JPG`,
      comment: "",
    },
    {
      id: 8,
      title: "Page 7",
      image: `/worldly/8.JPG`,
      comment: "",
    },
    {
      id: 9,
      title: "Concept",
      image: `/worldly/concept.JPG`,
      comment: "",
    },
  ];

  return (
    <div className="sketchOuterContainer">
      <div className="navBarDiv">
        <NavBar scrollText="STOP PHILOSOPHY. WHO CARES? TV IS GOD. SUBSTANCE ANXIETY. BREAKDOWN. STOP PHILOSOPHY. WHO CARES? TV IS GOD. SUBSTANCE ANXIETY. BREAKDOWN." />
      </div>
      <div className="sketchTitle">
        <h2>WORLDLY</h2>
        <p>
          WORLDLY is a zine I worked on at the end of 2023. It was a time I felt
          that I had been consumed with worries of career and social status. I
          noticed my own departure from philisophical and spiritual
          contemplation, and wanted to express the shallowness of materialism
          and commercialism. This was my (somewhat successful) attempt to
          reshift my focus to things that really matter.
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
export default Art_Worldly;
