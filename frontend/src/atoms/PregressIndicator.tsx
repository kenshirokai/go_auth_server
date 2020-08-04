import React from "react";
import styled, { keyframes } from "styled-components";

/**
 *@ Element & Styles
 */
const rotate = keyframes`
    from {
        transform: scale(0.5);
    }
    to {
        transform: scale(1);
    }
`
const Node = styled.span`
   display: inline-block;
   width: 100px;
   height: 100px;
   /* border: 2px solid #ebebeb; */
   border-radius: 100px;
   font-size: 3.3rem;
   color: #63cdda;
   /* animation: ${rotate} 5s linear infinite; */
`;

/**
 *@ ReactComponent
 */
const ProgressIndicator = () => {
    return (
        <Node>Loding...</Node>
    );
}

export default ProgressIndicator;