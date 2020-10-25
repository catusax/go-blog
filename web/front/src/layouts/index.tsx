import React from 'react';
import "./index.css"
import Header from './components/header'
import Footer from './components/footer'
import getinfo from '@/utils/siteinfo'

export default class extends React.Component {
  constructor(props: any) {
    super(props)
  }
  componentDidMount() {
    getinfo().then(siteinfo => {
      this.setState(siteinfo)
    })
  }
  state!: siteinfo;
  render() {
    return (
      <>
        <div className="wrap">
          <Header siteinfo={this.state} />
          {React.Children.map(this.props.children, (child: any) => {
            return React.cloneElement(child, { siteinfo: this.state })
          })}
          <Footer siteinfo={this.state} />
        </div>
      </>
    );
  }
}