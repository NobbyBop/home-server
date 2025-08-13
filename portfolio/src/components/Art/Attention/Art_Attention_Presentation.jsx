// import { SnavLink } from "react-router-dom";
import SnavLink from "../../SnavLink";
import NavBar from "../../NavBar";
function Art_Attention() {
  const body = document.body;
  body.className = "default-bg";
  const pageList = Array.from({ length: 23 }, (_, i) => (i + 1).toString());
  return (
    <div className="attentionOuterContainer">
      <div className="navBarDiv">
        <NavBar scrollText="ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION! ATTENTION!" />
      </div>
      <div className="attentionBody">
        <img
          className="attentionTitle"
          src={`/attention/title.png`}
        />
        <br></br>
        <SnavLink to="/art/attention/">BACK TO MAIN PAGE</SnavLink>
        <h2>PART THREE: THE USB</h2>
        <p>PLEASE FIND THE SLIDES OF THE PRESENTATION BELOW.</p>
        <div className="attentionPageList">
          {pageList.map((page) => (
            <img
              className="attentionPage"
              src={`/attention/presentation_slides/${page}.png`}
            />
          ))}
        </div>
        <SnavLink to="/art/attention/">BACK TO MAIN PAGE</SnavLink>
      </div>
    </div>
  );
}
export default Art_Attention;
