import React from "react";
import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const SwaggerUIComponent = () => {
  // Relaxed typing workaround: cast SwaggerUI to any to bypass TS errors
  const SwaggerUIAny = SwaggerUI as any;
  return <SwaggerUIAny url="/openapi.yaml" />;
};

export default SwaggerUIComponent;
