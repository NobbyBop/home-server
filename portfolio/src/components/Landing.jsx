import Navigation from "./Navigation";
import { useState } from "react";
function Landing() {
  const body = document.body;
  body.className = "landing-bg";
  let subtitles = [
    "developer. artist. quote collector.",
    "wilkomen, bienvenue, welcome.",
    "i would've stolen you a whole orchestra.",
    "see you space cowboy...",
    "the best jerry, the best!",
    "youths...",
    "that was a game. this is paintball.",
    "ain't that just the way?",
    "i'm gonna eat the whole thing.",
    "i think breakfast rocks, dude.",
    "pump up the jam.",
  ];
  let [subtitle, setSubtitle] = useState("developer. artist. quote collector.");
  let [subtitleIndex, setSubtitleIndex] = useState(0);
  let scrollTextA =
    "WELCOME. ARE YOU READY? THE TIME IS NOW. DON'T BE AFRAID. A WHOLE WORLD LIES AHEAD. GO ON.";
  let scrollTextB =
    "WE HOPE YOU ENJOY. YOU ARE ABOUT TO SEE SOMETHING NEW. TODAY IS GOING TO BE A GREAT DAY. GOOD MORNING.";

  const handleRefresh = function () {
    setSubtitleIndex((prevIndex) => {
      const newIndex = (prevIndex + 1) % subtitles.length;
      setSubtitle(subtitles[newIndex]);
      return newIndex;
    });
  };
  return (
    <div className="landingContainer">
      <h1 className="scrolling-text">
        <div className="scrolling-text-innerA">
          {scrollTextA}&nbsp;{scrollTextA}
        </div>
      </h1>
      <div className="landingContents">
        <h1 className="landingTitle">Nicholas Mirigliani</h1>
        <p className="landingSubtitle"> artist & developer
          {/* <button onClick={handleRefresh}>{subtitle}</button> */}
        </p>
        <Navigation />
      </div>
      <h1 className="scrolling-text">
        <div className="scrolling-text-innerB">
          {scrollTextB}&nbsp;{scrollTextB}
        </div>
      </h1>
    </div>
  );
}
export default Landing;
