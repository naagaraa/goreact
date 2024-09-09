import { Outlet, Link } from 'react-router-dom';

function About() {
  return (
    <>
      <p>about</p>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/about">About</Link>
        </li>
      </ul>
      <Outlet />
    </>
  )
}

export default About
