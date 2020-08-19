const dev =  {
    apiBaseUrl: `http://localhost:9001`,
    httpMode: `cors`
}

const prod = {
    apiBaseUrl: ``,
    httpMode: `same-origin`
}

const ENV = process.env.NODE_ENV == 'production' ? prod : dev;

export default ENV