import React, { useMemo } from "react";
import styled from "styled-components";

/**
 *@ Materials
 */
import Overlay from "../atoms/Overlay";
import ProgressIndicator from "../atoms/PregressIndicator";
import { transform } from "lodash";

/**
 *@ Elements & Styles
 */
const Node = styled.div`
    transition: 0.5s;
`;
const IndicatorWrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  bottom: 0;
  margin: 0 auto;
  transform: translateY(-50%);
  z-index: 11;
`;
/**
 *@ Types
 */
type Props = {
    isOpen: boolean
}

const Loading: React.FC<Props> = (props) => {
  const nodeStyles = useMemo(() => ({
      opacity: props.isOpen? `1` : `0`,
  }), [
      props.isOpen
  ])
  return (
    <Node style={nodeStyles}>
      <Overlay />
      <IndicatorWrapper>
        <ProgressIndicator />
      </IndicatorWrapper>
    </Node>
  );
};

export default Loading;
