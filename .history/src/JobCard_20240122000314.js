import React from 'react';

const JobCard = ({ job }) => {
  return (
    <div className="card h-100">
      <img src={job.companyLogo} className="card-img-top" alt={`${job.companyName} Logo`} />
      <div className="card-body">
        <h5 className="card-title">{job.jobTitle}</h5>
        <p className="card-text">{job.companyName}</p>
        <p className="card-text">{job.jobGeo}</p>
        <p className="card-text">{job.jobLevel}</p>
        <p className="card-text">{job.jobIndustry.join(', ')}</p>
        <p className="card-text">{job.jobType.join(', ')}</p>
        <p className="card-text">{job.pubDate}</p>
        {/* Additional job details here */}
      </div>
    </div>
  );
};

export default JobCard;
