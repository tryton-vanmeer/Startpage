use rand::prelude::*;
use sycamore::prelude::*;

fn random_color() -> String {
    vec![
        "flamingo", "pink", "mauve", "red", "maroon", "peach", "yellow", "green", "teal", "sky",
        "sapphire", "blue", "lavender",
    ]
    .choose_multiple(&mut rand::thread_rng(), 1)
    .cloned()
    .collect()
}

#[component(inline_props)]
fn Link<G: Html>(cx: Scope, name: String, url: String) -> View<G> {
    view! {cx,
        a(class=(format!("link {}", random_color())), href=(url)) {(name)}
    }
}

fn main() {
    sycamore::render(|cx| {
        view! { cx,
            Link(name="Steam".to_string(), url="https://store.steampowered.com".to_string())
        }
    });
}
