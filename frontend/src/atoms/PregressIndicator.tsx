import React from "react";
import styled, { keyframes } from "styled-components";

/**
 *@ Element & Styles
 */
const rotate = keyframes`
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);
    }
`
const Node = styled.span`
   display: inline-block;
   width: 100px;
   height: 100px;
   background-color: red;
   animation: ${rotate} 5s linear infinite;
`;

/**
 *@ ReactComponent
 */
const ProgressIndicator = () => {
    return (
        <Node></Node>
    );
}

export default ProgressIndicator;