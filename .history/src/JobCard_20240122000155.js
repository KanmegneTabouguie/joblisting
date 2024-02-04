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

        <h2 className="mt-4">Job Description:</h2>
        <div dangerouslySetInnerHTML={{ __html: job.jobDescription }} />

        <h2 className="mt-4">Qualifications:</h2>
        <div dangerouslySetInnerHTML={{ __html: job.qualifications }} />

        <h2 className="mt-4">Additional Information:</h2>
        <div dangerouslySetInnerHTML={{ __html: job.additionalInformation }} />
      </div>
    </div>
  );
};

export default JobCard;
