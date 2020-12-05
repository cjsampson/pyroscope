import React from 'react';
import { connect } from "react-redux";
import "react-dom"
import Spinner from "react-svg-spinner";
import DateRangePicker from "./DateRangePicker";
import DownloadButton from './DownloadButton';
import RefreshButton from './RefreshButton';
import SVGRenderer from "./SVGRenderer";
import LabelsFilter from "./LabelsFilter";
import Label from "./Label";
import NameSelector from "./NameSelector";

import classNames from "classnames";

class ProfileApp extends React.Component {
  constructor(props) {
    super(props);
  }

  renderURL() {
    let width = document.body.clientWidth - 30;
    let url = `/render?from=${encodeURIComponent(this.props.from)}&until=${encodeURIComponent(this.props.until)}&width=${width}`;
    let nameLabel = this.props.labels.find(x => x.name == "__name__");
    if (nameLabel) {
      url += "&name="+nameLabel.value+"{";
    } else {
      url += "&name=unknown{";
    }

    url += this.props.labels.filter(x => x.name != "__name__").map(x => `${x.name}=${x.value}`).join(",");
    url += "}";
    if(this.props.refreshToken){
      url += `&refreshToken=${this.props.refreshToken}`
    }
    return url;
  }

  render() {
    let renderURL = this.renderURL();
    return (
      <div className="todo-app">
        <div className="navbar">
          <h1 className="logo"></h1>
          <div className="labels">
            <NameSelector/>
            {this.props.labels.filter(x => x.name !== "__name__").map(function(label) {
              return <Label key={label.name} label={label}></Label>;
            })}
          </div>
          <LabelsFilter />
          <div className="navbar-space-filler"></div>
          <div className={
            classNames("navbar-spinner-container", {
              visible: this.props.isSVGLoading
            })
          }>
            <Spinner color="rgba(255,255,255,0.6)" size="20px"/>
          </div>
          <DownloadButton renderURL={renderURL+"&format=svg&download-filename=flamegraph.svg"} />
          &nbsp;
          <RefreshButton />
          &nbsp;
          <DateRangePicker />
        </div>
        <SVGRenderer renderURL={renderURL+"&format=frontend"}/>
      </div>
    );
  }
}


export default connect(
  (x) => x,
  {}
)(ProfileApp);