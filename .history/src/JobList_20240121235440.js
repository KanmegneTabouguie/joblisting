import React, { useState, useEffect } from 'react';
import JobCard from './JobCard';

const JobList = () => {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('https://jobicy.com/api/v2/remote-jobs');
        const data = await response.json();
        
        // Assuming the API response is an array of jobs
        if (Array.isArray(data)) {
          setJobs(data);
        } else {
          console.error('Invalid API response:', data);
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="container mt-5">
      <div className="row">
        {jobs.map(job => (
          <div key={job.id} className="col-md-6 mb-4">
            <JobCard job={job} />
          </div>
        ))}
      </div>
    </div>
  );
};

export default JobList;
