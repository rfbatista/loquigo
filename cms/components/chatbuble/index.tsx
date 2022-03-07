import styled from 'styled-components';

const Container = styled.div`
  p {
    // layout
    position: relative;
    max-width: 30em;
    color: #272f31;
    // looks
    background-color: #fff;
    padding: 8px 10px;
    font-size: 1.25em;
    border-radius: 0 4px 4px 4px;
    box-shadow: 0 0.125rem 0.5rem rgba(0, 0, 0, 0.3),
      0 0.0625rem 0.125rem rgba(0, 0, 0, 0.2);
  }

  p::before {
    // layout
    content: '';
    position: absolute;
    width: 0;
    height: 0;
    bottom: 100%;
    left: 10px; // offset should move with padding of parent
    top: 0em;
    border: 1rem solid transparent;
    border-top: none;
    border-left: none;
    transform: translate(-25px, 0px) rotate(180deg);
    box-sizing: border-box;
    //box-shadow: -3px 3px 3px 0 rgba(0, 0, 0, 0.4);
    // looks
    border-bottom-color: #fff;
    filter: drop-shadow(0 -0.0625rem 1.4625rem rgba(0.3, 0.3, 0.3, 0.1));
  }
`;

const ChatBubble = ({data}) => {
  return (
    <>
      <Container>
        <p>
          😔 Eu não te entendi. Envie a letra de uma das opções da lista para
          responder essa pergunta.{' '}
        </p>
      </Container>
    </>
  );
};

export default ChatBubble;
