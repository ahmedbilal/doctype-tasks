## Doctype Basic Tasks

### Go Task

1) Write a struct in Go for a Person object that has a numeric Age attribute.
    1) Write two functions in Go with the following function signature: `func ([]Person) bool`
    1) The first function should return `true` if there is a person who is exactly twice as old as any other person in the list, otherwise the function returns `false`.
    1) The second function should return `true` if there is a person who is at least twice as old as any other person in the list, otherwise the function returns `false`.

    ```golang
    type Person struct {
	    Age uint
    }

    func isTwiceOldAsSomeone(people []Person) bool {
        var peopleMap = make(map[uint]Person)

        for _, person := range people {
            peopleMap[person.Age] = person
            _, double_found := peopleMap[person.Age*2]

            // This default case caters for a person whose age is an odd number
            // because dividing an odd number by 2 results in a truncate operation
            var half_found = false

            if person.Age%2 == 0 {
                _, half_found = peopleMap[person.Age/2]
            }

            if double_found || half_found {
                return true
            }
        }
        return false
    }

    func isAtleastTwiceAsOldAsSomeone(people []Person) bool {
        // This function assumes that people array is not empty i.e it contain atleast one element
        var minMax = func(people []Person) (uint, uint) {
            var min = people[0].Age
            var max = people[0].Age
            for _, person := range people {
                if person.Age < min {
                    min = person.Age
                }
                if person.Age > max {
                    max = person.Age
                }
            }
            return min, max
        }
        var min, max = minMax(people)
        return max >= 2*min
    }
    ```
    4) Are you familiar with computational complexity? If so, what is the time complexity of the functions you implemented?

        **Yes. The time complexity of the first function would be `O(n)` where `n` is number of elements in the array/slice, as Map operations i.e insertion and lookup are `O(1)` but we are doing it for all elements.**

        **The time complexity for second function would be `O(n)` as we are traversing whole list only once and doing simple comparison and assignment which are considered `O(1)` operations.**

2) Write the DDL for creating the tables for a MariaDB database with users, tweets and followers.
    * A user has a username, password, email, first name and last name.
    * A tweet has text, a publication date and is posted by a user.
    * A user can be followed by other users.

    * The tables may contain additional fields that provide better space efficiency or performance.

    ```sql
    CREATE TABLE users (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(64) UNIQUE NOT NULL,
        password VARCHAR(64) NOT NULL,
        email VARCHAR(64) UNIQUE NOT NULL,
        first_name VARCHAR(64) NOT NULL,
        last_name VARCHAR(64) NOT NULL,

        UNIQUE KEY (username)
    )

    CREATE TABLE tweets (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        text VARCHAR(140) NOT NULL,
        publication_date DATETIME DEFAULT(NOW()) NOT NULL,
        user BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE
    )

    CREATE TABLE followers (
        user BIGINT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        follower BIGINT NOT NULL CHECK (follower <> user) REFERENCES users (id) ON DELETE CASCADE,

        PRIMARY KEY (user, follower)
    )

    ```

    1) Write a SQL query that returns the 30 latest tweets by users followed by the user with username "Mark". The result must include username, first name, last name, tweet text and publication date.

    ```sql
        SELECT username, first_name, last_name, text, publication_date FROM tweets
        INNER JOIN users ON tweets.user = users.id
        WHERE tweets.`user` IN (
            SELECT user
            FROM users
            INNER JOIN followers ON followers.follower = users.id
            WHERE users.username = "Mark"
        )

        ORDER BY publication_date DESC
        LIMIT 30
    ```

    2) One minute later, Mark is scrolling down the page to load 30 more tweets. What would the SQL query look like to fetch the next 30 tweets?

        **It would look the same except the `LIMIT 30` would be `LIMIT 30, 30` the first argument here is the offset.**

3) Write a React component that has a text field, a button and a 50x50 circle filled with a solid color. By default, the color of the circle is black. The user should be able to write a color in the text field and click the button to update the color of the circle.

    **You can play with it interactively at https://codepen.io/ahmedbilal/pen/vYGzXby?editors=0010**

    ```javascript
    class InteractiveCircle extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
        textboxValue: "black",
        color: "black",
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleClick = this.handleClick.bind(this);
    }
    handleClick(event) {
        this.setState({
        color: this.state.textboxValue
        });
        event.preventDefault();
    }
    handleChange(event) {
        this.setState({textboxValue: event.target.value});
    }

    render() {
        let circleStyle = {
        width: "50px",
        height: "50px",
        backgroundColor: this.state.color,
        borderRadius: "25px"
        }
        return (
        <div>
            <div id="circle" style={circleStyle}></div>
            <form style={{paddingTop: "10px"}}>
            <input type="text" value={this.state.textboxValue} onChange={this.handleChange} />
            <input type="submit" value="Change Color" onClick={this.handleClick} />
            </form>
        </div>
        );
    }
    }

    ReactDOM.render(
    <InteractiveCircle />,
    document.getElementById('root')
    )
    ```