// import { SnavLink } from "react-router-dom";
import SnavLink from "../../SnavLink";
import NavBar from "../../NavBar";
const body = document.body;
body.className = "default-bg";
function Art_Attention() {
  const pageList = Array.from({ length: 5 }, (_, i) => (i + 1).toString());
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
        <h2>PART TWO: THE BOOK</h2>
        <p>PLEASE FIND THE PARTICIPANTS' RESPONSES BELOW.</p>
        <div className="attentionPageList">
          {pageList.map((page) => (
            <img
              className="attentionPage"
              src={`/attention/participant${page}.JPEG`}
            />
          ))}
        </div>
        <SnavLink to="/art/attention/">BACK TO MAIN PAGE</SnavLink>
      </div>
    </div>
  );
}
export default Art_Attention;
