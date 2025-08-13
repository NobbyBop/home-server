import { NavLink, useNavigate } from "react-router-dom";

const SnavLink = ({ to, children, ...props }) => {
  const navigate = useNavigate();

  const handleClick = (e) => {
    e.preventDefault();
    navigate(to);
    window.scrollTo({
      top: 0,
      behavior: "auto",
    });
  };

  return (
    <NavLink to={to} onClick={handleClick} {...props}>
      {children}
    </NavLink>
  );
};

export default SnavLink;
