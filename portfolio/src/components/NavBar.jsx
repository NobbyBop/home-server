// import { SnavLink } from "react-router-dom";
import SnavLink from "./SnavLink";
import Navigation from "./Navigation";
function NavBar(props) {
  let scrollText;
  if (props.scrollText) scrollText = props.scrollText;
  else
    scrollText =
      "WHAT ARE YOU LOOKING FOR? WHAT WILL YOU SEE? WHAT WILL YOU THINK? WHAT WILL YOU FEEL? WHAT WILL YOU LIKE? WHAT WILL YOU HATE?";
  return (
    <div className="navContainer">
      <SnavLink className="headerNavLink" to="/">
        <h1>Nicholas Mirigliani</h1>
      </SnavLink>
      <Navigation />
      <p className="scrolling-textB">
        <div className="scrolling-text-innerC">
          {scrollText}&nbsp;{scrollText}&nbsp;{scrollText}&nbsp;{scrollText}
        </div>
      </p>
    </div>
  );
}
export default NavBar;
