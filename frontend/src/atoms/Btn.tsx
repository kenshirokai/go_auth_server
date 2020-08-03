import React, { useMemo, memo } from "react";
import styled from "styled-components";
import { merge } from "lodash";

/**
 *@ Elements & style
 */
const BaseButton = styled.button`
  padding: 10px 15px;
  font-size: 1.4rem;
  outline: none;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  transition: 0.5s;
  &:hover {
    opacity: 0.7;
  }
`;

/**
 *@ Types
 */
type Options = {
  backgroundColor: string;
  color: string;
};
type Props = {
  text: string;
  onClick: (event: React.MouseEvent<HTMLButtonElement>) => void;
} & Partial<Options>;

/**
 *@ default
 */
const defaultOptions = (): Options => ({
  backgroundColor: "#ebebeb",
  color: "#000000"
});

/**
 *@ ReactComponent
 */
const Btn: React.FC<Props> = memo((props) => {
  const styles = useMemo(
    () =>
      merge(defaultOptions(), {
        backgroundColor: props.backgroundColor,
        color: props.color
      }),
    [props.backgroundColor, props.color]
  );
  return (
    <BaseButton onClick={props.onClick} style={styles}>
      {props.text}
    </BaseButton>
  );
});

export default Btn;