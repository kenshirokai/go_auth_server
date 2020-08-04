import React from "react";
import styled from "styled-components";

/**
 *@ Element & Styles
 */
const Node = styled.div`
    position: fixed;
    width: 100vw;
    height: 100vh;
    background-color: rgba(255, 255, 255, 0.6);
    z-index: 10;
`;

/**
 *@ ReactComponent
 */
const Overlay = () => {
    return (
        <Node></Node>
    );
}

export default Overlay;