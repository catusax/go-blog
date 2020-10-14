
import React from 'react';
import { PluginComponent } from 'react-markdown-editor-lite';

class description extends PluginComponent{
  // 这里定义插件名称，注意不能重复
  static pluginName = 'counter';
  // 定义按钮被防止在哪个位置，默认为左侧，还可以放置在右侧（right）
  static align = 'left';
  constructor(props: any) {
    super(props);

    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    // 调用API，往编辑器中插入
    this.editor.insertText("<!--more-->");
  }

  render() {
    return (
      <span
        className="button button-type-counter"
        title="Counter"
        onClick={this.handleClick}
      >
        摘要
      </span>
    );
  }
}

export default description