// App.js
import React from 'react';
import Navbar from './Navbar';
import JobList from './JobList';
import Footer from './Footer'; // Import the Footer component
import CTAComponent from './CTAComponent';


function App() {
  return (
    <div className="App">
      {/* Include the Navbar component */}
      <Navbar />


      <main className="container">
        <CTAComponent />
        {/* Content of your App component */}
      
        <JobList />
      </main>
      <Footer /> 

    </div>
  );
}

export default App;
