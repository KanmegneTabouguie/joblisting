// Navbar.js
import React from 'react';

const Navbar = () => {
  return (
    <nav className="navbar">
      <div className="container">
        <a href="/" className="navbar-brand">
          Remote Jobs
        </a>

        {/* Additional navigation items */}
        <ul className="navbar-nav">
          <li className="nav-item">
            <a href="/work" className="nav-link">
              Work
            </a>
          </li>
          <li className="nav-item">
            <a href="/new" className="nav-link">
              New
            </a>
          </li>
          <li className="nav-item">
            <a href="/about" className="nav-link">
              About
            </a>
          </li>
          <li className="nav-item">
            <a href="/career" className="nav-link">
              Career
            </a>
          </li>
          <li className="nav-item">
            <a href="/contact" className="nav-link">
              Contact
            </a>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;
