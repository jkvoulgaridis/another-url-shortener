import React, { useState } from "react"
import "./Form.css"
const API_URL = 'http://127.0.0.1:8005'
const get_mappings="/get-mapping/"
const post_mapping = "/create-mapping/"
const go_to_mapping = "/go-to-mapping/"
export default class UrlForm extends React.Component {

	constructor(props) {
		super(props);
		this.state={
			url : "",
			shortUrl: ""
		}
		this.handleChange = this.handleChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}

	handleChange(event) {
		this.setState ({url : event.target.value})
		console.log(this.state.url)
	}

	handleSubmit(event){
		var myHeaders = new Headers();
		myHeaders.append("Content-Type", "application/json");

		var raw = JSON.stringify({
		  "url": this.state.url
		});

		var requestOptions = {
		  method: 'POST',
		  headers: myHeaders,
		  body: raw,
		  redirect: 'follow'
		};

		fetch("http://localhost:8005/create-mapping/", requestOptions)
		  .then(response => response.json())
		  .then(result => {
		  	console.log(result)
		  	this.setState({shortUrl: result["key"]})
		  })
		  .catch(error => console.log('error', error));
		  event.preventDefault();
	}

	render() {
	    return (
	      <form onSubmit={this.handleSubmit}>
	        <label>
	          What is your big URL?
	          <input type="text" value={this.state.url} onChange={this.handleChange} />
	        </label>

	        <input type="submit" value="Submit" />
	      
	       <div>
           <div>Here is your short Url:</div>
           <div>
            <a href={`${API_URL}${go_to_mapping}${this.state.shortUrl}`} target="_blank" rel="noopener noreferrer" className="short-url">
           	{`${API_URL}${go_to_mapping}${this.state.shortUrl}`}
           	</a>
           </div>
           </div>
	       	
	        
	      </form>
	    );
  	}
	
}