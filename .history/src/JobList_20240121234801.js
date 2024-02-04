import React, { useState, useEffect } from 'react';

const JobList = () => {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('https://jobicy.com/api/v2/remote-jobs');
        const data = await response.json();
        setJobs(data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      {/* Render job cards or list items here */}
    </div>
  );
};

export default JobList;
