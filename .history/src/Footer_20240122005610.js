import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

const Footer = () => {
  return (
    <footer className="mt-auto py-3 bg-danger">
      <Container>
        <Row className="d-flex justify-content-center">
          <Col md={4}>
            <span>Copyright &copy; 2024 Remote Jobs</span>
          </Col>
          <Col md={4} className="text-center">
            <span>Terms of Service</span>
          </Col>
          <Col md={4} className="text-end">
            <span>Privacy Policy</span>
          </Col>
        </Row>
      </Container>
    </footer>
  );
};

export default Footer;
