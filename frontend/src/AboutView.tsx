import { Component, createSignal, For, onMount } from 'solid-js'
import { IMetadata } from './interfaces';
import './AboutView.scss'

const BlogListView: Component = () => {
    const [blogPosts, setBlogPosts] = createSignal<IMetadata[]>([]);
    onMount(async () => {
        const res = await fetch(`http://localhost:3001/api/metadata`);
        setBlogPosts(await res.json());
      });
      
  return (
    <div id="about">
      <p>Jag har framförallt jobbat med webutveckling, men även en del applikationsutveckling till både desktop och mobil. Trivs bra när jag får frihet att välja vilka verktyg jag vill använda och kan påverka vad jag ska göra. Älskar att lära mig nya saker och att dela med mig av min kunskap.</p>

      <p>Under mina år som utvecklare har jag jobbat på projekt som exempelvis låssystem och projekthanteringsystem, men även med orderhantering och integrationer mot ekonomisystem.</p>

      <p>När jag inte jobbar så umgås jag mycket med min fru och mina döttrar. Min passion har varit att åka snowboard men efter flytten från norrland ner till östergötland har jag fått sadla om till klättring och discgolf.</p>
    </div>
  );
};

export default BlogListView;
