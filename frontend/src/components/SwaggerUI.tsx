import React from "react";
import * as SwaggerUI from "swagger-ui-react";
import SwaggerUIProps from "swagger-ui-react/swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

const SwaggerUIComponent = () => {
  const Swag = SwaggerUI as React.FunctionComponent<SwaggerUIProps>;
  return <Swag url="/openapi.yaml" />;
};

export default SwaggerUIComponent;
