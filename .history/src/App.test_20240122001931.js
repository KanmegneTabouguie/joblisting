// src/App.test.js
import { render, screen } from '@testing-library/react';
import App from './App';
import Modal from 'react-modal';

// Set the root element for accessibility
Modal.setAppElement(document.createElement('div'));

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
