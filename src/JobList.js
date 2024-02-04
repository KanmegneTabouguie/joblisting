import React, { useState, useEffect } from 'react';
import JobCard from './JobCard';

const JobList = () => {
  const [jobs, setJobs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('https://jobicy.com/api/v2/remote-jobs');
        const data = await response.json();

        console.log('API Response:', data);

        if (data && Array.isArray(data.jobs)) {
          setJobs(data.jobs);
        } else {
          console.error('Invalid API response:', data);
        }

        setLoading(false);
      } catch (error) {
        console.error('Error fetching data:', error);
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="container mt-5">
      <h1 className="text-center mb-4">Remote Job Listings</h1>

      {loading && <p className="text-center">Loading...</p>}

      <div className="row row-cols-1 row-cols-md-4">
        {jobs.map(job => (
          <div key={job.id} className="col mb-4">
            <JobCard job={job} />
          </div>
        ))}
      </div>
    </div>
  );
};

export default JobList;
