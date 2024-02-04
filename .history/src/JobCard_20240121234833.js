import React from 'react';

const JobCard = ({ job }) => {
  return (
    <div className="card">
      <div className="card-body">
        <h5 className="card-title">{job.title}</h5>
        <p className="card-text">{job.company_name}</p>
        {/* Add more job details here */}
      </div>
    </div>
  );
};

export default JobCard;
