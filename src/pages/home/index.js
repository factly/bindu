import React from "react";
import { Link } from 'react-router-dom';

function Home() {
  return (
    <div className="container">
      <Link to={'/templates'}>Templates</Link>
    </div>
  );
}

export default Home;