import React, { useState } from 'react';

const JobCard = ({ job }) => {
  const [showDetails, setShowDetails] = useState(false);

  const toggleDetails = () => {
    setShowDetails(!showDetails);
  };

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

        {/* Button to open modal */}
        <button className="btn btn-primary" onClick={toggleDetails}>
          Show Details
        </button>

        {/* Modal for additional job details */}
        {showDetails && (
          <div className="modal-overlay">
            <div className="modal">
              <button className="close-btn" onClick={toggleDetails}>&times;</button>
              <h2 className="mt-4">Job Description:</h2>
              <div dangerouslySetInnerHTML={{ __html: job.jobDescription }} />

              <h2 className="mt-4">Qualifications:</h2>
              <div dangerouslySetInnerHTML={{ __html: job.qualifications }} />

              <h2 className="mt-4">Additional Information:</h2>
              <div dangerouslySetInnerHTML={{ __html: job.additionalInformation }} />
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default JobCard;
