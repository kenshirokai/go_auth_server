import React from 'react'
import styled from 'styled-components';
import {HashRouter as Router, Switch, Route} from 'react-router-dom';

/**
 *@ Pages
 */
import Login from './pages/Login';
import ErrorPage from './pages/ErrorPage';

const Node = styled.div``

const App = () => {
    return (
        <React.Fragment>
            <Router>
                <Switch>
                    <Route exact path={'/'} component={Login}/>
                    <Route exact path={'/error'} component={ErrorPage}/>
                </Switch>
            </Router>
        </React.Fragment>
        
    );
}

export default App;