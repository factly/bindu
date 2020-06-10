import React, { useState } from "react";
import { Button } from "antd";

import Templates from "../templates/index.js";
function Home() {
  const [isNewClicked, setNewClicked] = useState(false);
  return (
    <div className="container">
      {!isNewClicked ? (
        <Button type="primary" onClick={() => setNewClicked(true)}>
          New Visualizations
        </Button>
      ) : (
        <Templates />
      )}
    </div>
  );
}

export default Home;
