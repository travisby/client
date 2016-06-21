/* @flow */
import React, {Component} from 'react'
import {Box, Avatar, Text} from '../common-adapters'
import TabBar, {TabBarItem} from '../common-adapters/tab-bar'
import {globalStyles, globalColors} from '../styles/style-guide'
import type {Props, UserInfo} from './friendships'

type UserEntryProps = UserInfo & {
  onClick?: (username: string) => void
};

const UserEntry = ({onClick, username, followsYou, following}: UserEntryProps) => (
  <Box style={userEntryContainerStyle} onClick={() => { onClick && onClick(username) }}>
    <Avatar style={userEntryAvatarStyle} size={48} username={username} followsYou={followsYou} following={following} />
    <Text type='BodySmall' style={userEntryUsernameStyle(followsYou)}>{username}</Text>
  </Box>
)

const userEntryContainerStyle = {
  ...globalStyles.clickable,
  ...globalStyles.flexBoxColumn,
  alignItems: 'center',
  justifyContent: 'flex-start',
  width: 130,
  padding: '10px 5px',
}

const userEntryAvatarStyle = {
  marginBottom: 2,
}

const userEntryUsernameStyle = followsYou => ({
  color: followsYou ? globalColors.green : globalColors.blue,
  textAlign: 'center',
})

class Render extends Component<void, Props, void> {
  render () {
    return (
      <TabBar style={this.props.style}>
        <TabBarItem
          selected={this.props.currentTab === 'Followers'}
          label={`FOLLOWERS (${this.props.followers.length})`}
          onClick={() => { this.props.onSwitchTab && this.props.onSwitchTab('Followers') }}>
          <Box style={tabItemContainerStyle}>
            {this.props.followers.map(user => <UserEntry key={user.username} {...user} onClick={this.props.onUserClick} />)}
          </Box>
        </TabBarItem>
        <TabBarItem
          selected={this.props.currentTab === 'Following'}
          label={`FOLLOWING (${this.props.following.length})`}
          onClick={() => { this.props.onSwitchTab && this.props.onSwitchTab('Following') }}>
          <Box style={tabItemContainerStyle}>
            {this.props.following.map(user => <UserEntry key={user.username} {...user} onClick={this.props.onUserClick} />)}
          </Box>
        </TabBarItem>
      </TabBar>
    )
  }
}

const tabItemContainerStyle = {
  ...globalStyles.flexBoxRow,
  flex: 1,
  flexWrap: 'wrap',
  borderTop: `solid 1px ${globalColors.black_10}`,
}

export default Render
