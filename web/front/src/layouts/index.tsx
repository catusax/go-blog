import React from 'react';
import "./index.css"
import Header from './components/header'
import Footer from './components/footer'
import getinfo from '@/utils/siteinfo'

export default function(props: { children: any; }) {
  getinfo().then(siteinfo=>{
    sessionStorage.setItem("SiteName",siteinfo.SiteName)
  })
  
    return (
      <>
      <div className="wrap">
          <Header prop={props}/>
          { props.children }
          <Footer/>
      </div>

      </>
    );
   }