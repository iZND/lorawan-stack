// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { Component } from 'react'
import bind from 'autobind-decorator'
import classnames from 'classnames'
import clipboard from 'clipboard'
import { defineMessages } from 'react-intl'

import Icon from '../icon'
import Message from '../../lib/components/message'
import PropTypes from '../../lib/prop-types'

import style from './safe-inspector.styl'

function chunkArray (array, chunkSize) {
  return Array.from(
    { length: Math.ceil(array.length / chunkSize) },
    (_, index) => array.slice(index * chunkSize, (index + 1) * chunkSize)
  )
}

function selectText (node) {
  if (document.body.createTextRange) {
    const range = document.body.createTextRange()
    range.moveToElementText(node)
    range.select()
  } else if (window.getSelection) {
    const selection = window.getSelection()
    const range = document.createRange()
    range.selectNodeContents(node)
    selection.removeAllRanges()
    selection.addRange(range)
  }
}

const m = defineMessages({
  copied: 'Copied to clipboard!',
  toggleVisibility: 'Toggle visibility',
  copyClipboard: 'Copy to clipboard',
  cStyle: 'Toggle C-Style formatting',
  byteSignificance: 'Swap byte significance',
})

@bind
export class SafeInspector extends Component {
  constructor (props) {
    super(props)

    this.state = {
      hidden: (props.hideable && !props.initiallyVisible) || false,
      byteStyle: true,
      copied: false,
      msb: true,
    }

    this.displayElem = React.createRef()
  }

  handleVisibiltyToggle () {
    this.setState(prev => ({ hidden: !prev.hidden }))
  }

  handleTransformToggle () {
    this.setState(prev => ({ byteStyle: !prev.byteStyle }))
  }

  handleSwapToggle () {
    this.setState(prev => ({ msb: !prev.msb }))
  }

  handleDataClick () {
    if (!this.state.hidden) {
      selectText(this.displayElem.current)
    }
  }

  handleCopyClick () {
    this.setState({ copied: true })
  }

  handleCopyAnimationEnd () {
    this.setState({ copied: false })
  }

  componentDidMount () {
    new clipboard('.copy')
  }

  render () {
    const {
      hidden,
      byteStyle,
      msb,
      copied,
    } = this.state

    const { data, isBytes, hideable } = this.props

    let formattedData = data
    let display = formattedData

    if (isBytes) {
      const chunks = chunkArray(data.toUpperCase().split(''), 2)
      if (!byteStyle) {
        const orderedChunks = msb ? chunks : chunks.reverse()
        formattedData = display = `{ ${orderedChunks.map(chunk => ` 0x${chunk.join('')}`)} }`
      } else {
        display = chunks.map((chunk, index) => (<span key={`${data}_chunk_${index}`}>{chunk}</span>))
      }
    }

    if (hidden) {
      display = '•'.repeat(formattedData.length)
    }

    return (
      <div className={style.container}>
        <div ref={this.displayElem} onClick={this.handleDataClick} className={style.data}>{display}</div>
        {!hidden && !byteStyle && isBytes && (
          <React.Fragment>
            <span>{ msb ? 'msb' : 'lsb' }</span>
            <button
              title={m.byteSignificance}
              className={style.buttonSwap}
              onClick={this.handleSwapToggle}
            >
              <Icon small icon="swap_horiz" />
            </button>
          </React.Fragment>
        )}
        {!hidden && isBytes && (
          <button
            title={m.cStyle}
            className={style.buttonTransform}
            onClick={this.handleTransformToggle}
          >
            <Icon small icon="code" />
          </button>
        )}
        <button
          title={m.copyClipboard}
          className={classnames(style.buttonCopy, 'copy')}
          onClick={this.handleCopyClick}
          data-clipboard-text={formattedData}
        >
          <Icon onClick={this.handleCopyClick} small icon="file_copy" />
          {copied && (
            <Message
              content={m.copied}
              onAnimationEnd={this.handleCopyAnimationEnd}
              className={style.copyConfirm}
            />
          )}
        </button>
        { hideable && (
          <button
            title={m.toggleVisibility}
            className={style.buttonVisibility}
            onClick={this.handleVisibiltyToggle}
          >
            <Icon small icon={hidden ? 'visibility' : 'visibility_off'} />
          </button>
        )}
      </div>
    )
  }
}

SafeInspector.defaultProps = {
  isBytes: true,
  initiallyVisible: false,
  hideable: true,
}

SafeInspector.propTypes = {
  /** The data to be displayed */
  data: PropTypes.string.isRequired,
  /** Whether the data is in byte format */
  isBytes: PropTypes.bool,
  /** Whether the data is initially visible */
  initiallyVisible: PropTypes.bool,
  /** Whether the data can be hidden (like passwords) */
  hideable: PropTypes.bool,
}

export default SafeInspector