import React from 'react'
import './pagination.css'

interface prop {
    onChange: onchange
    current: number
    pagesize: number
    total: number
}

interface onchange {
    (page: number): void
}

class Pagination extends React.Component<prop>{
    constructor(props: prop) {
        super(props)
    }
    state = {
        current: this.props.current || 1,
        pagesize: this.props.pagesize || 10,
        total: this.props.total
    }

    componentWillReceiveProps(props: { current: any; pagesize: any; total: any; }) {
        this.setState({
            current: props.current,
            pagesize: props.pagesize,
            total: props.total
        })
    }

    pagination(): any {
        const { current, pagesize, total } = this.state;
        let pages = []
        let totalpage = Math.ceil(total / pagesize)

        if (current > 1) { //有上一页
            pages.push(<a className="extend prev" onClick={() => {
                this.setState({ //PREV
                    current: current - 1
                })
                this.props.onChange(current - 1)
            }}>PREV</a>
            )

            if (current > 3) { //防止重复出现第一页
                pages.push(<a className="extend prev" onClick={() => {
                    this.setState({//第一页
                        current: 1
                    })
                    this.props.onChange(1)
                }}>1</a>
                )
            }

            if (current > 4) { //有省略号
                pages.push(
                    <span className="space">…</span>
                )
            }

            if (current > 2) { //上两页
                pages.push(<a className="extend prev" onClick={() => {
                    this.setState({
                        current: current - 2
                    })
                    this.props.onChange(current - 2)
                }}>{current - 2}</a>
                )
            }

            pages.push(<a className="extend prev" onClick={() => {
                this.setState({ //上一页
                    current: current - 1
                })
                this.props.onChange(current - 1)
            }}>{current - 1}</a>
            )
        }

        pages.push(//当前页
            <span className="page-number current">{current}</span>
        )

        if (current < totalpage) { //有下一页

            pages.push(<a className="extend prev" onClick={() => {
                this.setState({ //下一页
                    current: current + 1
                })
                this.props.onChange(current + 1)
            }}>{current + 1}</a>
            )

            if (current < totalpage - 1) { //有后两页
                pages.push(<a className="extend prev" onClick={() => {
                    this.setState({ //下两页
                        current: current + 2
                    })
                    this.props.onChange(current + 2)
                }}>{current + 2}</a>
                )
            }

            if (current < totalpage - 3) {//省略号
                pages.push(
                    <span className="space">…</span>
                )
            }

            if (current < totalpage - 2) { //最后一页
                pages.push(<a className="extend prev" onClick={() => {
                    this.setState({
                        current: totalpage
                    })
                    this.props.onChange(totalpage)
                }}>{totalpage}</a>
                )
            }

            pages.push(<a className="extend prev" onClick={() => {
                this.setState({ //NEXT
                    current: current + 1
                })
                this.props.onChange(current + 1)
            }}>NEXT</a>
            )
        }

        return pages
    }

    render() {
        return (
            <nav id="pagination">
                <div className="pagination">
                    {this.pagination()}
                </div>
            </nav>
        )
    }
}
export default Pagination