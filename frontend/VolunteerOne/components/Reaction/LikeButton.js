
import React from "react";
import { Block, Button, Text } from "galio-framework";
import { StyleSheet, Image } from "react-native";
import { Svg } from "react-native-svg";
import Icon from "../Icon";

class LikeButton extends React.Component {
    state = {
        count: 0
    }

    render() {    
        const incrementLikes = () => {
            console.log("Like button pressed")
            let newCount = this.state.count + 1
            this.setState ({
                count: newCount
            })
            
        }

        return (
            <Button style={styles.button}
            onPress={incrementLikes} >      
      <Block flex row>
                <Icon
      family="MaterialIcons"
      size={20}
      name="thumb-up"
      color="#32325D"
    />
            <Text style={styles.titleText}>Likes: {this.state.count}</Text>
            </Block>
          </Button>

        );
        }
        
}
const styles = StyleSheet.create({
    titleText: {
        fontSize: 18,
        color: "#32325D",
        marginLeft: 10,
      },
    button: {
        borderless: true,
        backgroundColor: '#fff',
        shadowColor: '#fff',
        width: 'auto',
        height: 20,
        marginTop: -30,
        marginLeft: 20,
    }

});

export default LikeButton;
    