import React, { lazy, Suspense } from "react";
import styled from "styled-components";
import { HashRouter as Router, Switch, Route } from "react-router-dom";

/**
 *@ Pages
 */
import Login from "./pages/Login";
import SignUp from './pages/SignUp'

//import ErrorPage from './pages/ErrorPage';
const ErrorPage = lazy(() => import("./pages/ErrorPage"));

const Node = styled.div``;

const App = () => {
  return (
    <React.Fragment>
      <Suspense fallback={<div>loading...</div>}>
        <Router>
          <Switch>
            <Route exact path={"/login"} component={Login} />
            <Route exact path={"/signup"} component={SignUp} />
            <Route exact path={"/error"} component={ErrorPage} />
          </Switch>
        </Router>
      </Suspense>
    </React.Fragment>
  );
};

export default App;
