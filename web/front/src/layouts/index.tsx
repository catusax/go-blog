import React from 'react';
import "./index.css"
import Header from './components/header'
import Footer from './components/footer'

export default function(props: { children: any; }) {

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