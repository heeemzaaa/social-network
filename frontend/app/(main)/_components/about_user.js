export default function AboutUser({aboutMe}) {
    return (
        <section className="about_me_container p2">
            <span>About me</span>
            {aboutMe && <p>{aboutMe}</p>}
        </section>
    );
}