import React, { ChangeEvent, useMemo } from "react";
import styled from "styled-components";
import { merge } from "lodash";

/************************
 *ラベル付きInput
 ************************/

/**
 *@ Materials
 */
import Label from "../atoms/Label";
import Input from "../atoms/Input";

/**
 *@ Elements & Styles
 */
const Node = styled.div`
  display: flex;
  align-items: center;
  width: 100%;
`;
const LabelWrapper = styled.div`
  width: 100px;
  margin-right: 20px;
`;
/**
 *@ Type
 */
type Options = {
  onChange: ((e: ChangeEvent<HTMLInputElement>) => void) | undefined;
};
type Props = {
  labelText: string;
} & Partial<Options>;
/**
 *@ default
 */
const defaultOptions = (): Options => ({
  onChange: undefined,
});
/**
 *@ ReactComponent
 */
const LabelInput: React.FC<Props> = (props) => {
  const options = useMemo(
    () =>
      merge(defaultOptions(), {
        onChange: props.onChange,
      }),
    [props.onChange]
  );
  return (
    <Node>
      <LabelWrapper>
        <Label value={props.labelText} />
      </LabelWrapper>
      <Input onChange={options.onChange} />
    </Node>
  );
};

export default LabelInput;
