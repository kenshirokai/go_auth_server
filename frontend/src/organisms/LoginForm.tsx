import React, { useMemo, ChangeEvent } from "react";
import styled from "styled-components";
import { merge } from "lodash";

/**
 *@ Materials
 */
import LabelInput from "../molecules/LabelInput";
import Btn from "../atoms/Btn";

/**
 *@ Element & Styles
 */
const Node = styled.div`
  padding: 20px;
  padding-top: 40px;
  border: 1px solid #ebebeb;
  border-radius: 10px;
`;
const InputWrapper = styled.div`
  width: 100%;
  margin-bottom: 20px;
`;
const ButtonWrapper = styled.div`
  display: flex;
  justify-content: flex-end;
  width: 100%;
`;
/**
 *@ Types
 */
type Options = {
  emailChange: ((e: ChangeEvent<HTMLInputElement>) => void) | undefined;
  passwordChange: ((e: ChangeEvent<HTMLInputElement>) => void) | undefined;
};
type Props = {
  onClick: () => void;
} & Partial<Options>;
/**
 *@ default
 */
const defaultOptions = (): Options => ({
  emailChange: undefined,
  passwordChange: undefined,
});
/**
 *@ ReactComponent
 */
const LoginForm: React.FC<Props> = (props) => {
  const options = useMemo(
    () =>
      merge(defaultOptions(), {
        emailChange: props.emailChange,
        passwordChange: props.passwordChange,
      }),
    [props.emailChange, props.passwordChange]
  );
  return (
    <Node>
      <InputWrapper>
        <LabelInput labelText={"email"} onChange={options.emailChange} />
      </InputWrapper>
      <InputWrapper>
        <LabelInput labelText={"password"} onChange={options.passwordChange} />
      </InputWrapper>
      <ButtonWrapper>
        <Btn text={"Login"} onClick={props.onClick} />
      </ButtonWrapper>
    </Node>
  );
};
export default LoginForm;
