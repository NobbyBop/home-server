import { NavLink } from "react-router-dom";
import NavBar from "./NavBar";
function InProgress() {
  return (
    <>
      <NavBar scrollText="COMING SOON. PLEASE BE PATIENT. IT WILL BE WELL WORTH THE WAIT. HOLD OUT A LITTLE LONGER. COMING SOON. PLEASE BE PATIENT. IT WILL BE WELL WORTH THE WAIT. HOLD OUT A LITTLE LONGER." />
      <div className="wipDiv">
        <p>In progress!</p>
      </div>
    </>
  );
}
export default InProgress;
