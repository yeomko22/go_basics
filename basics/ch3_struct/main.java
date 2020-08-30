class Person {
    String name;
    int age;

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
    void info(){
        System.out.printf("name: %s, age: %d\n", this.name, this.age);
    }
    void introduce() {
        System.out.print("Hello, this is my info: ");
        this.info();
    }
}

class Player extends Person {
    String team;
    public Player(String name, int age, String team) {
        super(name, age);
        this.team = team;
    }
    void info() {
        System.out.printf("name: %s, age: %d, team: %s\n", this.name, this.age, this.team);
    }
}

public class main {
    public static void SayName(Person p) {
        System.out.printf("Hello %s\n", p.name);
    }
    public static void main(String[] args) {
        Person person = new Person("John", 30);
        Player player = new Player("Lilard", 25, "portland");
        person.info();
        player.info();
        person.introduce();
        player.introduce();
        SayName(person);
        SayName(player);
    }
}
