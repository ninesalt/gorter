let envURL = process.env.REACT_APP_API_URL
let config = {
    "moduleURL": envURL ? envURL : 'http://localhost:5000',
    "codes": [200, 300, 304],
    "browsers": ["Chrome", "Firefox", "Safari", "Opera", "Internet Explorer"],
    "upstreamModules": ["something", "somethingelse"],
}

// API Endpoints on server
config.getAllRulesAPI = config.moduleURL + '/getRules'          //GET
config.addNewRuleAPI = config.moduleURL + "/addNew"    //POST
config.deleteRuleAPI = config.moduleURL + '/remove'       //POST

export default config