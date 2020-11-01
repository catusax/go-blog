import React from 'react';
import {Link} from 'umi'
import style from './404.less'
export default class extends React.Component{

    render(){
        return(
            <div className={style.center}>
                <h2>未找到页面</h2>
                <Link className={style.btn} to="/"><span>回到主页</span></Link>
            </div>
        )
    }
}