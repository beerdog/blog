import { Component, createResource, createSignal, For, onMount, splitProps } from 'solid-js'
import { IMetadata } from './interfaces';
import './BlogView.scss'
import { useParams } from 'solid-app-router';

const BlogListView: Component = (props) => {
  const params = useParams();
  const fetchBlogPost = async (id: string) => (await fetch(`http://localhost:3001/api/blogposts/${id}`)).text()
  const [blogPost] = createResource(params.id, fetchBlogPost);
  // onMount(async () => {
  //     const res = await fetch(`http://localhost:3001/api/blogposts/${params.id}`);
  //     setBlogPost(await res.text());
  //   });
      
  return (
    <div id="blog-post">
      <div innerHTML={blogPost()}></div>
    </div>
  );
};

export default BlogListView;
