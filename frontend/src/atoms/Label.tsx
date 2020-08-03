import React, { memo } from "react";
import styled from "styled-components";

/**
 *@ Elements & Styles
 */
const Text = styled.span`
    font-size: 1.4rem;
`;

/**
 *@ Types
 */
type Props = {
  value: string;
};
/**
 *@ ReactComponent
 */
const Label: React.FC<Props> = memo((props) => {
  return <Text>{props.value}</Text>;
});

export default Label;
