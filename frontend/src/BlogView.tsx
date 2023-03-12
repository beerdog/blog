import { Component, createResource, createSignal, For, onMount, splitProps } from 'solid-js'
import { IMetadata } from './interfaces';
import './BlogView.scss'
import { useParams } from 'solid-app-router';

const BlogListView: Component = (props) => {
  const params = useParams();
  const fetchBlogPost = async (id: string) => (await fetch(`http://localhost:3001/api/blogposts/${id}`)).text()
  const fetchMetadata = async (id: string) => (await fetch(`http://localhost:3001/api/metadata/${id}`)).json() as unknown as IMetadata
  const [blogPost] = createResource(params.id, fetchBlogPost);
  const [metadata] = createResource(params.id, fetchMetadata);
  
  return (
    <div id="blog-post">
      <a class="back" href="/">Back to blog posts</a>
      <h1>{metadata()?.title}</h1>
      <div class="preamble">{metadata()?.preamble}</div>
      <div class="date-category">Posted at {metadata()?.publishDate} in {metadata()?.category}</div>
      <div class="tags">{metadata()?.tags.map(t => `#${t} `)}</div>
      <div innerHTML={blogPost()}></div>
    </div>
  );
};

export default BlogListView;
