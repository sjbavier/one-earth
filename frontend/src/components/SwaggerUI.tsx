import "swagger-ui-react/swagger-ui.css";
import SwaggerUI from "swagger-ui-react";

const SwaggerUIComponent = () => {
  const Swag = SwaggerUI as any;
  return <Swag url="/openapi.yaml" />;
};

export default SwaggerUIComponent;
