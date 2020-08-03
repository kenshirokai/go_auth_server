import React, { useMemo } from "react";
import { useHistory } from "react-router-dom";
import styled from "styled-components";

/**
 *@ Materials
 */
import LoginForm from "../organisms/LoginForm";
import Loading from "../organisms/Loading";

/**
 *@ Hooks
 */
import usePost from "../hooks/usePost";
import useInput from "../hooks/useInput";

/**
 *@ Element & Styles
 */
const Node = styled.div`
  position: relative;
  height: 100vh;
`;
const FormWrapper = styled.div`
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  bottom: 0;
  max-width: 768px;
  width: 100%;
  margin-right: auto;
  margin-left: auto;
  transform: translateY(-50%);
`;

/**
 *@ ReactComponents
 */
const Login: React.FC = () => {
  const history = useHistory();
  const email = useInput();
  const password = useInput();
  const body = useMemo(
    () => ({
      email: email.state,
      password: password.state,
    }),
    [email.state, password.state]
  );
  const handleFailure = () => {
    history.push("/error");
  };
  const { loading, httpPost } = usePost<{ msg: string }>(
    "http://localhost:9001/users",
    body,
    { failure: handleFailure }
  );

  return (
    <Node>
      <Loading isOpen={loading} />
      {!loading && (
        <FormWrapper>
          <LoginForm
            emailChange={email.onChange}
            passwordChange={password.onChange}
            onClick={httpPost}
          />
        </FormWrapper>
      )}
    </Node>
  );
};

export default Login;
