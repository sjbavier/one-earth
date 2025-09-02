import React from "react";
import * as SwaggerUI from "swagger-ui-react";
// import SwaggerUIProps from "swagger-ui-react/swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const SwaggerUIComponent = () => {
  // const Swag = SwaggerUI as unknown as React.FunctionComponent<SwaggerUIProps>;
  const Swag = SwaggerUI as any;
  return <Swag url="/openapi.yaml" />;
};

export default SwaggerUIComponent;
