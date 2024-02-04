// App.js
import React from 'react';
import Navbar from './Navbar';
import JobList from './JobList';

function App() {
  return (
    <div className="App">
      {/* Include the Navbar component */}
      <Navbar />
      

      <main className="container">
        {/* Content of your App component */}
      
        <JobList />
      </main>
    </div>
  );
}

export default App;
